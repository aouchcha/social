import Link from "next/link";
import { useState } from "react"

export function CreateGroupComp() {
    const [name, setname] = useState("")
    const [desc, setdesc] = useState("")

    const handleChange = (event, helper) => {
        if (helper == "name") {
            setname(event.target.value);
        }else {
            setdesc(event.target.value)
        }
    };

    return (
        <div className="creategroup" style={{padding : "4%", display :"flex", flexDirection : "column", justifyContent : "center", textAlign :"center"}}>
            <h3>Groupe Name</h3>
            <input type="text" placeholder="Enter group name" onChange={(e) => handleChange(e,"name")} value={name}></input>
            <h3>Group Description</h3>
            <input type="text" placeholder="Enter group description" onChange={(e) =>handleChange(e,"desc")} value={desc}></input>
            <br></br>
            <button onClick={ async () => await CreateGroup({"user_id":4,"name":name, "description":desc})} style={{width : "fit-content"}}>Submit</button>
        </div>
    )
}



async function CreateGroup(body) {
    const res = await fetch("http://localhost:8080/api/groupe/create",{
        method : "post",
        body : JSON.stringify(body)
    })
    
    if (!res.ok) {
        const err = await res.json()
        console.log(err.msg);
    }
    window.location.href = "/groups"
}

export function CreateGroupButt() {
    return (
        <Link href={`/groups/groupcreation`} style={{padding:"10px", border:"1px solid red", backgroundColor:"red"}}>Create Groupe</Link>
    )
}