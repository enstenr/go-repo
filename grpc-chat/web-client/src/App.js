import { useState } from "react";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";

// Import your generated schemas
import { UserService, ChatService } from "./gen/pb/chat/v1/chat_pb";

// 1. Define Transports
const userTransport = createConnectTransport({
    baseUrl: "http://localhost:50051",
});

const chatTransport = createConnectTransport({
    baseUrl: "http://localhost:50052",
});

// 2. Create Clients
const userClient = createClient(UserService, userTransport);
const chatClient = createClient(ChatService, chatTransport);
function App() {
    const [myId, setMyId] = useState(""); // Your User ID
    const [targetId, setTargetId] = useState(""); // Who you want to message
    const [isConnected, setIsConnected] = useState(false);
    const [messages, setMessages] = useState([]);
    const [inputText, setInputText] = useState("");

    const startChat = async () => {
        if (!myId) return alert("Please enter your User ID first!");

        setIsConnected(true);
        try {
            const stream = chatClient.subscribe({ userId: myId });
            console.log(`Subscribed as ${myId}`);

            for await (const msg of stream) {
                setMessages((prev) => [...prev, { from: msg.senderId, text: msg.text }]);
            }
        } catch (err) {
            console.error("Stream closed:", err);
            setIsConnected(false);
        }
    };

    const sendMessage = async (e) => {
        e.preventDefault();
        await chatClient.sendMessage({
            senderId: myId,
            recipientId: targetId || myId, // If target is empty, echo to self
            text: inputText,
        });
        setInputText("");
    };

    return (
        <div style={{ padding: "20px" }}>
            <h2>Connect RPC Chat</h2>

            {!isConnected ? (
                <div>
                    <input placeholder="Your ID (e.g. user-1)" value={myId} onChange={e => setMyId(e.target.value)} />
                    <button onClick={startChat}>Login & Listen</button>
                </div>
            ) : (
                <div>
                    <p style={{color: "green"}}>● Connected as: <strong>{myId}</strong></p>
                    <input placeholder="Recipient ID" value={targetId} onChange={e => setTargetId(e.target.value)} />

                    <div style={{ border: "1px solid #ccc", height: "200px", margin: "10px 0", overflowY: "auto", padding: "10px" }}>
                        {messages.map((m, i) => (
                            <div key={i}><strong>{m.from}:</strong> {m.text}</div>
                        ))}
                    </div>

                    <form onSubmit={sendMessage}>
                        <input value={inputText} onChange={e => setInputText(e.target.value)} placeholder="Message..." />
                        <button type="submit">Send</button>
                    </form>
                </div>
            )}
        </div>
    );
}
export default App;