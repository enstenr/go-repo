import React, { useState } from "react";
import { Toaster } from "react-hot-toast";
import { useChat } from "./hooks/useChat";
import { LoginView } from "./components/LoginView";
import { Sidebar } from "./components/Sidebar";
import { ChatWindow } from "./components/ChatWindow";
import { User } from "./gen/pb/chat/v1/chat_pb";
import { TechHeader } from "./components/TechHeader";

const App: React.FC = () => {
    const { myUser, users, messages, register, sendMessage, fetchUsers } = useChat();
    const [targetUser, setTargetUser] = useState<User | null>(null);

    return (
        <div style={{ display: "flex", flexDirection: "column", height: "100vh", width: "100vw", overflow: "hidden" }}>
            <Toaster position="top-center" />

            {/* 1. Static Header */}
            <TechHeader myUser={myUser} />

            {/* 2. Dynamic Content Area */}
            <div style={{ display: "flex", flex: 1, overflow: "hidden" }}>
                {!myUser ? (
                    /* If no user, center the LoginView in the remaining space */
                    <div style={{ flex: 1, display: "flex", justifyContent: "center", alignItems: "center", backgroundColor: "#f9fafb" }}>
                        <LoginView onRegister={register} />
                    </div>
                ) : (
                    /* If user exists, show the dashboard */
                    <>
                        <Sidebar
                            users={users}
                            myUser={myUser}
                            activeUser={targetUser}
                            onSelect={setTargetUser}
                            onRefresh={fetchUsers}
                        />
                        <ChatWindow
                            targetUser={targetUser}
                            myUser={myUser}
                            messages={messages}
                            onSendMessage={(text) => sendMessage(myUser.id, targetUser!.id, text)}
                        />
                    </>
                )}
            </div>
        </div>
    );
};

export default App;