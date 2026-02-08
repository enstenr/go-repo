import React from 'react';
import ReactDOM from 'react-dom/client';
import App from "./App";

// We find the element and tell TypeScript it exists using '!'
// or by doing a null check.
const rootElement = document.getElementById('root');

if (!rootElement) {
    throw new Error("Failed to find the root element");
}

const root = ReactDOM.createRoot(rootElement);

root.render(
    <React.StrictMode>
       <App />
    </React.StrictMode>
);