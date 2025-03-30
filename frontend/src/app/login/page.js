"use client";

import { useState } from "react";
import styles from "./login.module.css";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  const setCookie = (name, value, days = 7) => {
    const expires = new Date();
    expires.setTime(expires.getTime() + days * 24 * 60 * 60 * 1000);
    document.cookie = `${name}=${value};expires=${expires.toUTCString()};path=/`;
  };

  const handleLogin = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8080/api/signin", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });

      const res = await response.json();

      if (response.ok) {
        setCookie("token", res.token.Token);
        localStorage.setItem("user",  res.userId);
        setMessage("Login successful!");
      } else {
        throw new Error(res.msg);
      }
    } catch (err) {
      console.error("Error:", err);
      setMessage("Login failed. Check credentials.");
    }
  };

  return (
    <div className={styles.container}>
      <h2 className={styles.title}>Login</h2>
      {message && <p className={styles.message}>{message}</p>}
      <form onSubmit={handleLogin} className={styles.form}>
        <div className={styles.inputGroup}>
          <label>Email:</label>
          <input type="email" name="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
        </div>

        <div className={styles.inputGroup}>
          <label>Password:</label>
          <input type="password" name="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
        </div>

        <button type="submit" className={styles.button}>Login</button>
      </form>
    </div>
  );
}