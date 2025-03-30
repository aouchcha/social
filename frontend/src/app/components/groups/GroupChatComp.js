// app/components/ChatBox.js
import { useEffect, useState, useRef } from "react";

export function ChatBox({ chat, groupname, setchat, groupid, userid }) {
  // console.log({ chat: chat, check: Array.isArray(chat) });

  //   const [messages, setMessages] = useState([]);
  const [input, setInput] = useState("");
  const [ws, setWs] = useState(null);
  const chatRef = useRef(null);
  //   const userId = Math.floor(1);
  // console.log(userid);
  

  useEffect(() => {
    const socket = new WebSocket(`ws://localhost:8080/api/groupe/chat?groupid=${groupid}&&userid=${userid}`);
    socket.onopen = () => console.log("Connected to WebSocket");
    socket.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      console.log("msg,",msg);
      
      setchat((prev) => [...prev, msg]);
      chatRef.current?.scrollIntoView({ behavior: "smooth" });
    };
    socket.onclose = () => console.log("WebSocket Disconnected");
    setWs(socket);
    return () => socket.close();
  }, []);

  const sendMessage = () => {
    if (ws && input.trim() !== "") {
      const message = { "message": input };
      ws.send(JSON.stringify(message));
      // setchat((prev) => [...prev, message]);
      setInput("");
    }
  };
  

  return (
    <div className="chat-box" style={{ border: "1px solid black" }}>
      <div className="messages" id="chatScroll">
        {!chat || chat.length === 0 ? (
          <p>No messages to see</p>
        ) : (
          chat.map((content) => (
            <div key={content.message_id} className={`message ${content.user_id == userid ? "sent" : "received"}`}>
              <p className="sub">By : {content.username}</p><br></br>
              {content.message}
            </div>
          ))
        )}

        <div ref={chatRef}></div>
      </div>
      <div className="input-area">
        <input value={input} onChange={(e) => setInput(e.target.value)} placeholder="Type a message..." />
        <button onClick={sendMessage}>Send</button>
      </div>
      <style jsx>{`
        .chat-box {
          width: 400px;
          height: 500px;
          display: flex;
          flex-direction: column;
          border: 1px solid #ccc;
          background: white;
          border-radius: 10px;
          overflow: hidden;
        }
        .messages {
          flex-grow: 1;
          padding: 10px;
          overflow-y: auto;
          display: flex;
          flex-direction: column;
        }
        .message {
          padding: 10px;
          margin: 5px;
          border-radius: 8px;
          max-width: 70%;
        }
        .sent {
          align-self: flex-end;
          background-color: #007bff;
          color: white;
        }
        .received {
          align-self: flex-start;
          background-color: #e0e0e0;
          color: black;
        }
        .input-area {
          display: flex;
          padding: 10px;
          border-top: 1px solid #ccc;
        }
        .input-area input {
          flex-grow: 1;
          padding: 10px;
          border: 1px solid #ccc;
          border-radius: 5px;
          margin-right: 5px;
        }
        .input-area button {
          padding: 10px;
          background: #007bff;
          color: white;
          border: none;
          border-radius: 5px;
          cursor: pointer;
        }
        .sub {
          font-weight: lighter;
          font-size: smaller;
  font: optional;
  font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;

          color : red;
        }
      `}</style>
    </div>
  );
};
