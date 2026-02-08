export const theme = {
    colors: {
        primary: "#4f46e5",
        primaryLight: "#e0e7ff",
        bg: "#f9fafb",
        sidebar: "#f3f4f6",
        border: "#e5e7eb",
        textMain: "#111827",
        textMuted: "#6b7280",
        white: "#ffffff",
    },
    radius: {
        sm: "4px",
        md: "8px",
        lg: "12px",
        full: "9999px",
    },
    // Add to chatStyles in theme.ts

};

export const chatStyles: { [key: string]: React.CSSProperties } = {
    // ... your existing appLayout, bubble, actionButton ...

    // The main container for the chat interface
    chatArea: {
        flex: 1,
        display: "flex",
        flexDirection: "column",
        backgroundColor: theme.colors.white
    },

    // Top bar with the user's name
    chatHeader: {
        padding: "20px",
        borderBottom: `1px solid ${theme.colors.border}`,
        fontSize: "18px",
        fontWeight: "bold",
        color: theme.colors.textMain
    },

    // The scrollable area for messages
    messageBox: {
        flex: 1,
        padding: "20px",
        overflowY: "auto",
        display: "flex",
        flexDirection: "column",
        gap: "12px"
    },

    // Bottom container for the input and button
    inputArea: {
        padding: "20px",
        borderTop: `1px solid ${theme.colors.border}`,
        display: "flex",
        gap: "10px"
    },

    // The text input field
    chatInput: {
        flex: 1,
        padding: "12px 20px",
        borderRadius: theme.radius.full,
        border: `1px solid ${theme.colors.border}`,
        outline: "none",
        fontSize: "14px"
    },

    // Send button (specific variant of actionButton)
    sendButton: {
        padding: "0 24px",
        backgroundColor: theme.colors.primary,
        color: theme.colors.white,
        border: "none",
        borderRadius: theme.radius.full,
        cursor: "pointer",
        fontWeight: "600"
    },
    headerBanner: {
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
        padding: "0 25px",
        height: "64px", // Fixed height is often easier for layouts
        background: "linear-gradient(90deg, #4f46e5 0%, #7c3aed 100%)",
        color: "white",
        boxShadow: "0 2px 4px rgba(0,0,0,0.1)",
        boxSizing: "border-box",
        flexShrink: 0, // Prevents header from squishing
    },
    badgeContainer: {
        display: "flex",
        gap: "8px",
    },
    badge: {
        padding: "4px 10px",
        backgroundColor: "rgba(255, 255, 255, 0.2)",
        borderRadius: "4px",
        fontSize: "11px",
        fontWeight: "600",
        textTransform: "uppercase" as const,
        letterSpacing: "0.5px",
        border: "1px solid rgba(255, 255, 255, 0.3)"
    }
};