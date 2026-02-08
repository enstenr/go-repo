import { useState, useEffect, useCallback } from "react";
import { createClient, ConnectError } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { UserService, ChatService } from "../gen/pb/chat/v1/chat_pb";
import { User, Message } from "../gen/pb/chat/v1/chat_pb";
import toast from "react-hot-toast";

// Transports & Clients
const userClient = createClient(UserService, createConnectTransport({ baseUrl: "http://localhost:50051" }));
const chatClient = createClient(ChatService, createConnectTransport({ baseUrl: "http://localhost:50052" }));

export interface DisplayMessage {
    from: string;
    text: string;
}

export const useChat = () => {
    const [myUser, setMyUser] = useState<User | null>(null);
    const [users, setUsers] = useState<User[]>([]);
    const [messages, setMessages] = useState<DisplayMessage[]>([]);
    const [isListening, setIsListening] = useState(false);

    const fetchUsers = useCallback(async () => {
        try {
            const response = await userClient.listUsers({});
            setUsers(response.users);
        } catch (err) {
            console.error("Failed to fetch users:", err);
        }
    }, []);

    const startListening = useCallback(async (userId: string) => {
        if (isListening) return;
        setIsListening(true);
        try {
            const stream = chatClient.subscribe({ userId });
            for await (const msg of stream) {
                const incoming = msg as Message;
                setMessages((prev) => [...prev, { from: incoming.senderId, text: incoming.text }]);
            }
        } catch (err) {
            console.error("Stream error:", err);
            setIsListening(false);
        }
    }, [isListening]);

    const register = async (userName: string) => {
        try {
            const response = await userClient.register({ name: userName });
            if (response.user) {
                setMyUser(response.user);
                fetchUsers();
                startListening(response.user.id);
            }
        } catch (err) {
            toast.error("Registration failed. Check if Go server is running.");
        }
    };

    const sendMessage = async (senderId: string, recipientId: string, text: string) => {
        try {
            await chatClient.sendMessage({ senderId, recipientId, text });
        } catch (err) {
            if (err instanceof ConnectError) {
                toast.error(err.rawMessage);
            } else {
                toast.error("An unexpected error occurred.");
            }
        }
    };

    return { myUser, users, messages, register, sendMessage, fetchUsers };
};