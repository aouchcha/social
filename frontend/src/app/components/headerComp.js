import Link from 'next/link'
export function Header() {

    return (
        <>
            <div className="Header" style={{padding:"10px"}}>
                <Link href={`/`} style={{padding:"10px", border:"1px solid red", backgroundColor:"wheat"}}>Home</Link>
                <Link href={`/groups`} style={{padding:"10px", border:"1px solid red", backgroundColor:"wheat"}}>Groups</Link>
                <Link href={`/profile`} style={{padding:"10px", border:"1px solid red", backgroundColor:"wheat"}}>Profile</Link>
            </div>
        </>
    )
}   