"use client"

import { useParams } from "next/navigation"
import { useState, useEffect } from "react"
import Link from 'next/link'
import { ChatBox } from "./GroupChatComp"


export function SingleGroupComp() {
    const [events, setevents] = useState([])
    const [posts, setposts] = useState([])
    const [chat, setchat] = useState([])
    const [activeTab, setActiveTab] = useState("posts")
    const [GroupId, setGroupId] = useState(localStorage.getItem("currentGroupId"))
    const [UserId, setUserId] = useState(localStorage.getItem("user"))
    const params = useParams()
    const decodedGroupName = decodeURIComponent(params.groupname)

    async function GetGroupEvents() {
        try {
            const res = await fetch(`http://localhost:8080/api/events?groupid=${GroupId}`, {
                method: "get",
            })

            if (!res.ok) {
                setevents([])
                throw new Error("Error in fetching group events");
            }

            const data = await res.json()
            console.log(data);

            setevents(data.Events)
            setActiveTab("events")
        } catch (error) {
            console.log(error);
            setevents([])
        }
    }

    async function GetGroupPosts() {
        try {
            const res = await fetch(`http://localhost:8080/api/groupe/posts?groupid=${GroupId}`, {
                method: "get",
            })

            if (!res.ok) {
                setposts([])
                throw new Error("Error in fetching group posts")
            }

            const data = await res.json()
            console.log(data);

            setposts(data.Posts)
            setActiveTab("posts")
        } catch (error) {
            console.error("Failed to fetch posts:", error)
            setposts([])
        }
    }

    async function GetGroupChat() {
        try {
            const res = await fetch(`http://localhost:8080/api/groupe/load_messages?groupid=${GroupId}`, {
                method: "get"
            })
            if (!res.ok) {
                const err = res.json()
                throw new Error(err.msg);
            }
            const data = await res.json()
            console.log("chat" ,data.Groups);

            setchat(data.Groups)
            setActiveTab("chat")
        } catch (error) {
            console.log(error);
            setchat([])
        }
    }

    async function UpdateEvent(body) {
        console.log(body);

        try {
            const res = await fetch("http://localhost:8080/api/event/update", {
                method: "post",
                body: JSON.stringify(body)
            })
            if (!res.ok) {
                const err = await res.json()
                console.log(err.msg);

                throw new Error("Can't interact with the event");
            }
            console.log("WE update the event without any problems");
        } catch (error) {
            console.log(error);
        }
    }

    useEffect(() => {
        setGroupId(localStorage.getItem("currentGroupId"))
        setUserId(localStorage.getItem("user"))

        if (!params || !params.groupname) return
        (async function () {
            await GetGroupPosts()
        })();
    }, [params])

    return (
        <>
            <h1>{decodedGroupName}</h1>
            <div className="nav" style={{ display: "flex", gap: "20px" }}>
                <button onClick={GetGroupPosts} style={{ fontWeight: activeTab === "posts" ? "bold" : "normal" }}>Posts</button>
                <button onClick={GetGroupEvents} style={{ fontWeight: activeTab === "events" ? "bold" : "normal" }}>Events</button>
                <button onClick={GetGroupChat} style={{ fontWeight: activeTab === "chat" ? "bold" : "normal" }}>Chat</button>

            </div>
            {activeTab === "posts" &&
                <div className="posts">
                    {!posts || posts.length === 0 ? (
                        <p>There is no events to see</p>
                    ) : (
                        posts.map((post) => (
                            <div key={post.post_id} style={{ padding: "10px" }}>
                                <h2>Post Content : {post.content}</h2>
                                <p>Created By : {post.username}</p>
                                <sub>Created At : {post.created_at}</sub>
                                <br></br>
                                <sub><span>Comments Count : {post.comments_count}</span> ; <span>Likes Count : {post.likes}</span></sub>
                            </div>
                        ))
                    )}
                </div>
            }
            {activeTab === "events" &&
                <div className="events">
                    <Link href={`/groups/eventcreation`} style={{ padding: "10px", border: "1px solid red", backgroundColor: "red" }}>Create Event</Link>
                    {!events || events.length === 0 ? (
                        <p>There is no events to see</p>
                    ) : (
                        events.map((event) => (
                            <div key={event.event_id} style={{ padding: "10px", display: "flex", flexDirection: "column" }}>
                                <h2>Title : {event.title}</h2>
                                <h3><span>Date : {event.event_date}</span>// <span>Event Time : {event.event_time} </span></h3>
                                <p>Description : {event.description}</p>
                                <sub>Created By : {event.event_creator}</sub>
                                <br></br>
                                <sub><span>Created At : {event.created_at} </span></sub>
                                <button onClick={() => UpdateEvent({ "event_id": event.event_id, "user_id": event.user_id, "status": "g" })} style={{ width: "fit-content" }}>Going</button>
                                <button onClick={() => UpdateEvent({ "event_id": event.event_id, "user_id": event.user_id, "status": "n" })} style={{ width: "fit-content" }}>Not Going</button>
                            </div>
                        ))
                    )}
                </div>
            }
            {
                activeTab === "chat" &&
                <div className="groupchat">
                       <ChatBox chat={chat} groupname={params.groupname} setchat={setchat} groupid={GroupId} userid={UserId}/>
                </div>
            }
        </>
    )
} 
