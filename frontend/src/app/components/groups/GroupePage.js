import { useState, useEffect } from "react";

export function GroupsPage() {
    const [groups, setgroups] = useState("")
    useEffect(() => {
        async function fetchGroups() {
            try {
                const res = await fetch("http://localhost:8080/api/groups", {
                    method: "get"
                });
                
                if (!res.ok) {
                    const err = await res.json();
                   throw new Error(err.msg);
                }
                
                // Need to await the json() call
                const data = await res.json();
                console.log(data);
                
                setgroups(data.Groups);
            } catch (error) {
                console.log(error)
            }
        }
        fetchGroups();
    },[]);
    
    return (
            <div className="container">
                {!groups || groups.length === 0 ? (
                    <p>No Groups To See !!!</p>
                ):(
                    groups.map((group) => (
                        <div 
                        key={group.GroupId}  
                        style={{
                            padding: "2%", 
                            border: "1px solid black",
                            cursor: "pointer",
                            marginBottom: "10px",
                            borderRadius: "4px",
                            textAlign : "center"
                        }}>
                            <h1>Groupe Name : {group.Name}</h1>
                            <h2>Created By : {group.Created_By}</h2>
                            <p>Descreption : {group.Description}</p>
                            <button  onClick={() => SeeGroupContent(group.GroupId, group.IsMember, group.Name)} >{group.IsMember ? "See The Content" : "Send a request to join"}</button>
                        </div>
                    ))
                )}
            </div>
    );
}

async function SeeGroupContent(group_id, isMember, group_name) {
    if (!isMember) {
        const res = await fetch("http://localhost:8080/api/groupe/request", {
            method : "post",
            body : JSON.stringify({
                "group_id" : group_id,
                "requester_id" : 5,
                "status" : "p"
            })
        })
        if (!res.ok) {
            const err = await res.json()
            console.log(err.msg);
        }
    }else {
        localStorage.setItem('currentGroupId', group_id);
        window.location.href = `/groups/${group_name}`
    }
}