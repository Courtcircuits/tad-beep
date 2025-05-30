import { useParams } from '@tanstack/react-router'
import { useContext, useState } from 'react'
import { ChannelContext } from '../context/channel'
export function MessageInput() {
  const { sendMessage } = useContext(ChannelContext)
  const [message, setMessage] = useState('')
  const { channelId } = useParams({ from: '/channels/$channelId' })
  return (
    <div className="flex flex-row w-full gap-2">
      <input type="text" className="w-10/12" placeholder="Message" onChange={(e) => setMessage(e.target.value)} />
      <button className="w-2/12" onClick={() => sendMessage({
        content: message,
        owner: localStorage.getItem('name') || '',
        channel: channelId
      })}>Send</button>
    </div>
  )
}
