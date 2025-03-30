"use client"

import { CreateGroupButt } from "../components/groups/CreateGroup.js"
import {GroupsPage} from "../components/groups/GroupePage.js"
import { Header } from "../components/headerComp.js"

export default function GG() {
    return (
    <div style={{display:"flex", flexDirection:"column", gap :"10px"}}>
        <Header/>
        <CreateGroupButt/>
        <GroupsPage/>
    </div>
    )
}