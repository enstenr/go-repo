import React, { useState, useEffect, useCallback, FormEvent } from "react";
import {ConnectError, createClient} from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

// Import your generated schemas and types
// Note: Ensure your generated files are in the paths below
import { UserService, ChatService } from "./gen/pb/chat/v1/chat_pb";
import { User, Message } from "./gen/pb/chat/v1/chat_pb";
import toast, {Toaster} from "react-hot-toast";


// 1. Transports
const userTransport = createConnectTransport({
    baseUrl: "http://localhost:50051",
});

const chatTransport = createConnectTransport({
    baseUrl: "http://localhost:50052",
});

// 2. Clients
const userClient = createClient(UserService, userTransport);
const chatClient = createClient(ChatService, chatTransport);

// Define a simple local interface for message display logic
interface DisplayMessage {
    from: string;
    text: string;
}

const App: React.FC = () => {
    const [userName, setUserName] = useState<string>("");
    const [myUser, setMyUser] = useState<User | null>(null);
    const [users, setUsers] = useState<User[]>([]);
    const [targetUser, setTargetUser] = useState<User | null>(null);
    const [messages, setMessages] = useState<DisplayMessage[]>([]);
    const [inputText, setInputText] = useState<string>("");
    const [isListening, setIsListening] = useState<boolean>(false);

    // Fetch the list of all available users
    const fetchUsers = useCallback(async () => {
        try {
            const response = await userClient.listUsers({});
            setUsers(response.users);
        } catch (err) {
            console.error("Failed to fetch users:", err);
        }
    }, []);

    // Initial load
    useEffect(() => {
        fetchUsers();
    }, [fetchUsers]);

    // Start listening for incoming messages (Server Streaming)
    const startListening = async (userId: string) => {
        if (isListening) return;
        setIsListening(true);
        try {
            const stream = chatClient.subscribe({ userId: userId });
            console.log("Listening for messages as:", userId);

            for await (const msg of stream) {
                // Asserting msg type as Message if needed by your generator
                const incoming = msg as Message;
                setMessages((prev) => [...prev, { from: incoming.senderId, text: incoming.text }]);
            }
        } catch (err) {
            console.error("Stream error:", err);
            setIsListening(false);
        }
    };

    // Handle User Registration
    const handleRegister = async (e: FormEvent) => {
        e.preventDefault();
        if (!userName) return alert("Please enter a name");
        try {
            const response = await userClient.register({ name: userName });
            if (response.user) {
                setMyUser(response.user);
                await fetchUsers();
                startListening(response.user.id);
            }
        } catch (err) {
            alert("Registration failed. Is the Go server on 50051 running?");
        }
    };

    // Handle Sending Message
    const sendMessage = async (e: FormEvent) => {
        e.preventDefault();
        if (!inputText || !targetUser || !myUser) return;

        try {
            await chatClient.sendMessage({
                senderId: myUser.id,
                recipientId: targetUser.id,
                text: inputText,
            });
            // We don't add "Me" manually here anymore because your Go server
            // sends it back to the sender's stream!
            setInputText("");
        } catch (err) {
            console.error("Send failed:", err);
            if (err instanceof ConnectError) {
                // This pulls the "User is offline" message you wrote in Go
                toast.error(err.rawMessage);
            } else {
                toast.error("An unexpected error occurred.");
                console.error("Send failed:", err);
            }
        }
    };

    if (!myUser) {
        return (
            <div style={styles.centeredContainer}>
                <div style={styles.card}>
                    <h2>Join Chat</h2>
                    <form onSubmit={handleRegister}>
                        <input
                            style={styles.input}
                            placeholder="Enter your name..."
                            value={userName}
                            onChange={(e) => setUserName(e.target.value)}
                        />
                        <button style={styles.button} type="submit">Enter Room</button>
                    </form>
                </div>
            </div>
        );
    }

    return (
        <div style={styles.appContainer}>
            <Toaster position="top-center" reverseOrder={false} />
            <div style={styles.sidebar}>
                <h3 style={{ borderBottom: "1px solid #ddd", paddingBottom: "10px" }}>Users</h3>
                <button onClick={fetchUsers} style={styles.refreshBtn}>Refresh List</button>
                {users.map((u) => (
                    <div
                        key={u.id}
                        onClick={() => setTargetUser(u)}
                        style={{
                            ...styles.userItem,
                            backgroundColor: targetUser?.id === u.id ? "#e0e7ff" : "white",
                            fontWeight: u.id === myUser.id ? "bold" : "normal"
                        } as React.CSSProperties}
                    >
                        <span style={{ color: u.id === myUser.id ? "#4f46e5" : "#333" }}>
                            {u.name} {u.id === myUser.id ? "(You)" : ""}
                        </span>
                    </div>
                ))}
            </div>

            <div style={styles.chatArea}>
                <div style={styles.chatHeader}>
                    {targetUser ? `Chatting with ${targetUser.name}` : "Select a user to start chatting"}
                </div>

                <div style={styles.messageBox}>
                    {messages
                        .filter(m => !targetUser || m.from === targetUser.id || m.from === myUser.id)
                        .map((m, i) => {
                            const isMe = m.from === myUser.id;
                            return (
                                <div key={i} style={{
                                    ...styles.messageBubble,
                                    alignSelf: isMe ? "flex-end" : "flex-start",
                                    backgroundColor: isMe ? "#4f46e5" : "#f3f4f6",
                                    color: isMe ? "white" : "black",
                                } as React.CSSProperties}>
                                    <strong>{isMe ? "Me: " : ""}</strong>{m.text}
                                </div>
                            );
                        })}
                </div>

                {targetUser && (
                    <form onSubmit={sendMessage} style={styles.inputArea}>
                        <input
                            style={styles.chatInput}
                            value={inputText}
                            onChange={(e) => setInputText(e.target.value)}
                            placeholder={`Message ${targetUser.name}...`}
                        />
                        <button style={styles.sendButton} type="submit">Send</button>
                    </form>
                )}
            </div>
        </div>
    );
}

// Inline Styles with TypeScript support
const styles: { [key: string]: React.CSSProperties } = {
    centeredContainer: { display: "flex", justifyContent: "center", alignItems: "center", height: "100vh", backgroundColor: "#f9fafb" },
    card: { padding: "40px", backgroundColor: "white", borderRadius: "8px", boxShadow: "0 4px 6px rgba(0,0,0,0.1)", textAlign: "center" },
    input: { padding: "10px", width: "200px", marginRight: "10px", borderRadius: "4px", border: "1px solid #ccc" },
    button: { padding: "10px 20px", backgroundColor: "#4f46e5", color: "white", border: "none", borderRadius: "4px", cursor: "pointer" },
    appContainer: { display: "flex", height: "100vh", fontFamily: "sans-serif" },
    sidebar: { width: "260px", backgroundColor: "#f3f4f6", borderRight: "1px solid #ddd", padding: "20px", overflowY: "auto" },
    refreshBtn: { width: "100%", marginBottom: "10px", cursor: "pointer", fontSize: "12px" },
    userItem: { padding: "12px", cursor: "pointer", marginBottom: "5px", borderRadius: "6px", transition: "0.2s" },
    chatArea: { flex: 1, display: "flex", flexDirection: "column", backgroundColor: "white" },
    chatHeader: { padding: "20px", borderBottom: "1px solid #eee", fontSize: "18px", fontWeight: "bold" },
    messageBox: { flex: 1, padding: "20px", overflowY: "auto", display: "flex", flexDirection: "column", gap: "10px" },
    messageBubble: { padding: "10px 15px", borderRadius: "18px", maxWidth: "70%", fontSize: "14px" },
    inputArea: { padding: "20px", borderTop: "1px solid #eee", display: "flex", gap: "10px" },
    chatInput: { flex: 1, padding: "12px", borderRadius: "25px", border: "1px solid #ddd", outline: "none" },
    sendButton: { padding: "0 20px", backgroundColor: "#4f46e5", color: "white", border: "none", borderRadius: "25px", cursor: "pointer" }
};

export default App;