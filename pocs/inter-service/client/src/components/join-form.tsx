import { useNavigate } from '@tanstack/react-router'
import { useState } from 'react'

export function JoinForm() {
  const [name, setName] = useState('')
  const [channel, setChannel] = useState('')
  const navigate = useNavigate({from :'/'})

  const handleJoin = () => {
    localStorage.setItem('name', name)
    navigate({
      to: `/channels/${channel}`
    })
    console.log('joining channel', channel)
  }

  return (
    <div className="flex flex-col gap-2">
      <form className="flex flex-col gap-2" onSubmit={(e) => {
        handleJoin()
        e.preventDefault()
      }}>
        <h3>Join a channel</h3>
        <input
          type="text"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <input
          type="text"
          placeholder="Channel"
          value={channel}
          onChange={(e) => setChannel(e.target.value)}
        />
        <button type="submit">Join</button>
      </form>
    </div>
  )
}
