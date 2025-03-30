"use client";

import { useState } from "react";

import styles from "./register.module.css";

export default function Register() {

    const [username, setUsername] = useState("");
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [birthDate, setBirthDate] = useState("");
    const [avatar, setAvatar] = useState("");
    const [message, setMessage] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await fetch("http://localhost:8080/api/signup", {
                method: "POST",
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ username, firstName, lastName, email, password, birthDate, avatar })
            });

            const res = await response.json();
            if (response.ok) {
                setMessage("Registration successful!");
            } else {
                throw new Error(res.msg || "Registration failed.");
            }
        } catch (error) {
            console.error(error); // Fixed error variable
            setMessage("Registration failed. Check credentials.");
        }
    };

    return (
        <div className={styles.container}>
            <h2>Register User</h2>
            {message && <p className={styles.message}>{message}</p>} {/* Display message */}
            <form onSubmit={handleSubmit} className={styles.form}>
                <div className={styles.inputGroup}>
                    <label>Username:</label>
                    <input
                        type="text"
                        name="username"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                </div>

                <div className={styles.inputGroup}>
                    <label>First Name:</label>
                    <input
                        type="text"
                        name="firstName"
                        value={firstName}
                        onChange={(e) => setFirstName(e.target.value)}
                        required
                    />
                </div>

                <div className={styles.inputGroup}>
                    <label>Last Name:</label>
                    <input
                        type="text"
                        name="last_name"
                        value={lastName}
                        onChange={(e) => setLastName(e.target.value)}
                        required
                    />
                </div>

                <div className={styles.inputGroup}>
                    <label>Email:</label>
                    <input
                        type="email"
                        name="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                </div>

                <div className={styles.inputGroup}>
                    <label>Password:</label>
                    <input
                        type="password"
                        name="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>

                <div className={styles.inputGroup}>
                    <label>Date of Birth:</label>
                    <input
                        type="date"
                        name="date_of_birth"
                        value={birthDate}
                        onChange={(e) => setBirthDate(e.target.value)}
                        required
                    />
                </div>

                <div className={styles.inputGroup}>
                    <label>Avatar (optional):</label>
                    <input
                        type="file"
                        name="avatar"
                        // value={firstName}
                        onChange={(e) => setAvatar(e.target.value)} />
                </div>

                <button type="submit" className={styles.button}>
                    Register
                </button>
            </form>
        </div>
    );
}
