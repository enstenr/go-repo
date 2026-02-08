import React from "react";
import { User } from "../gen/pb/chat/v1/chat_pb";

interface SidebarProps {
    users: User[];
    myUser: User;
    activeUser: User | null;
    onSelect: (user: User) => void;
    onRefresh: () => void;
}

export const Sidebar: React.FC<SidebarProps> = ({
                                                    users,
                                                    myUser,
                                                    activeUser,
                                                    onSelect,
                                                    onRefresh
                                                }) => {
    return (
        <div style={styles.sidebar}>
            <div style={styles.header}>
                <h3 style={{ margin: 0 }}>Users</h3>
                <button onClick={onRefresh} style={styles.refreshBtn}>
                    Refresh
                </button>
            </div>

            <div style={styles.list}>
                {users.map((u) => {
                    const isMe = u.id === myUser.id;
                    const isActive = activeUser?.id === u.id;

                    return (
                        <div
                            key={u.id}
                            onClick={() => onSelect(u)}
                            style={{
                                ...styles.userItem,
                                backgroundColor: isActive ? "#e0e7ff" : "transparent",
                                fontWeight: isMe ? "bold" : "normal",
                                border: isActive ? "1px solid #c7d2fe" : "1px solid transparent"
                            }}
                        >
              <span style={{ color: isMe ? "#4f46e5" : "#374151" }}>
                {u.name} {isMe ? "(You)" : ""}
              </span>
                        </div>
                    );
                })}
            </div>
        </div>
    );
};

const styles: { [key: string]: React.CSSProperties } = {
    sidebar: {
        width: "280px",
        backgroundColor: "#f3f4f6",
        borderRight: "1px solid #e5e7eb",
        display: "flex",
        flexDirection: "column",
        height: "100%",
    },
    header: {
        padding: "20px",
        borderBottom: "1px solid #e5e7eb",
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center"
    },
    refreshBtn: {
        padding: "4px 8px",
        fontSize: "12px",
        cursor: "pointer",
        borderRadius: "4px",
        border: "1px solid #d1d5db"
    },
    list: {
        flex: 1,
        overflowY: "auto",
        padding: "10px"
    },
    userItem: {
        padding: "12px",
        cursor: "pointer",
        marginBottom: "4px",
        borderRadius: "8px",
        transition: "all 0.2s ease"
    },
};