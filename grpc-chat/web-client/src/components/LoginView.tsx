import React, { useState } from "react";
import {TechHeader} from "./TechHeader";

interface LoginViewProps {
    onRegister: (name: string) => void;
}

export const LoginView: React.FC<LoginViewProps> = ({ onRegister }) => {
    const [userName, setUserName] = useState("");

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (userName.trim()) {
            onRegister(userName);
        } else {
            alert("Please enter a name");
        }
    };

    return (
        <div style={styles.centeredContainer}>

            <div style={styles.card}>
                <h2 style={{ marginBottom: "20px", color: "#111827" }}>Join Chat</h2>
                <form onSubmit={handleSubmit}>
                    <input
                        style={styles.input}
                        placeholder="Enter your name..."
                        value={userName}
                        onChange={(e) => setUserName(e.target.value)}
                    />
                    <button style={styles.button} type="submit">
                        Enter Room
                    </button>
                </form>
            </div>
        </div>
    );
};

const styles: { [key: string]: React.CSSProperties } = {
    centeredContainer: {
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
        backgroundColor: "#f9fafb"
    },
    card: {
        padding: "40px",
        backgroundColor: "white",
        borderRadius: "12px",
        boxShadow: "0 10px 15px -3px rgba(0,0,0,0.1)",
        textAlign: "center"
    },
    input: {
        padding: "12px",
        width: "240px",
        marginRight: "10px",
        borderRadius: "6px",
        border: "1px solid #d1d5db",
        outline: "none"
    },
    button: {
        padding: "12px 24px",
        backgroundColor: "#4f46e5",
        color: "white",
        border: "none",
        borderRadius: "6px",
        cursor: "pointer",
        fontWeight: "600"
    },
};