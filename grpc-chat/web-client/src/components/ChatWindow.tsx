import React, { useState } from "react";
import { User } from "../gen/pb/chat/v1/chat_pb";
import { DisplayMessage } from "../hooks/useChat";
import { chatStyles as styles, theme } from "../styles/theme";
interface ChatWindowProps {
    targetUser: User | null;
    myUser: User;
    messages: DisplayMessage[];
    onSendMessage: (text: string) => void;
}

export const ChatWindow: React.FC<ChatWindowProps> = ({ targetUser, myUser, messages, onSendMessage }) => {
    const [inputText, setInputText] = useState("");

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (!inputText.trim()) return;
        onSendMessage(inputText);
        setInputText("");
    };

    const filteredMessages = messages.filter(
        (m) => !targetUser || m.from === targetUser.id || m.from === myUser.id
    );

    return (
        <div style={styles.chatArea}>
            <div style={styles.chatHeader}>
                {targetUser ? `Chatting with ${targetUser.name}` : "Select a user to start"}
            </div>
            <div style={styles.messageBox}>
                {filteredMessages.map((m, i) => {
                    const isMe = m.from === myUser.id;
                    return (
                        <div key={i} style={{
                            ...styles.bubble,
                            alignSelf: isMe ? "flex-end" : "flex-start",
                            backgroundColor: isMe ? "#4f46e5" : "#f3f4f6",
                            color: isMe ? "white" : "black"
                        }}>
                            <strong>{isMe ? "Me: " : ""}</strong>{m.text}
                        </div>
                    );
                })}
            </div>
            {targetUser && (
                <form onSubmit={handleSubmit} style={styles.inputArea}>
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
    );
};
