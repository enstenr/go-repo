import React from "react";
import { chatStyles as styles } from "../styles/theme";
import { User } from "../gen/pb/chat/v1/chat_pb";

interface TechHeaderProps {
    myUser: User | null;
}

export const TechHeader: React.FC<TechHeaderProps> = ({ myUser }) => {
    return (
        <header style={styles.headerBanner}>
            <div style={{ display: "flex", alignItems: "center", gap: "15px" }}>
                <h1 style={{ fontSize: "20px", margin: 0, fontWeight: "800" }}>
                    ConnectRPC Messenger
                </h1>
                <div style={styles.badgeContainer}>
                    <span style={styles.badge}>React 18</span>
                    <span style={styles.badge}>TypeScript</span>
                    <span style={styles.badge}>Buf / gRPC</span>
                    <span style={styles.badge}>Go Backend</span>
                </div>
            </div>

            <div style={{ fontSize: "13px", opacity: 0.9 }}>
                {myUser ? (
                    <>Logged in as: <strong>{myUser.name}</strong></>
                ) : (
                    <span style={{ fontStyle: "italic" }}>Awaiting Connection...</span>
                )}
            </div>
        </header>
    );
};