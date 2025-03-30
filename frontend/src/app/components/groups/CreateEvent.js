"use client"

import { useState } from "react"

export function CreateEventComp() {

    const [Title, setTitle] = useState("")
    const [Desc, setDesc] = useState("")
    const [Date, setDate] = useState("")
    const [Time, setTime] = useState("")

    const group_id = parseInt(localStorage.getItem("currentGroupId"))

    async function CreateEvent(body) {
        try {
            const res = await fetch("http://localhost:8080/api/event/create",{
                method : "post",
                body : JSON.stringify(body)
            })
            if (!res.ok) {
                const err = await res.json()
                throw new Error(err.msg);
            }
            console.log(await res.json());
            
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <>
            <div className="eventcreation" style={{display:"flex", flexDirection:"column", gap:"10px"}}>
                <h3>Event Title</h3>
                <input type="text" value={Title} onChange={(e) => setTitle(e.target.value)}></input>
                <h3>Event Description</h3>
                <input type="text" value={Desc} onChange={(e) => setDesc(e.target.value)}></input>
                <h3>Event Date</h3>
                <input type="date" value={Date} onChange={(e) => setDate(e.target.value)}></input>
                <h3>Event Time</h3>
                <input type="time" value={Time} onChange={(e) => setTime(e.target.value)}></input>
                <button 
                onClick={() => CreateEvent({
                    //this will be removed because the use id with be taking from context
                    "user_id" : 1,
                    "group_id" : group_id,
                    "title" : Title,
                    "description" : Desc,
                    "event_date" : Date,
                    "event_time" : Time
                })}
                >Submit</button>
            </div>
        </>
    )
}