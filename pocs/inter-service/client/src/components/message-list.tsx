import { useContext } from 'react'
import { ChannelContext, type Message } from '../context/channel'

export function MessageList() {
  const { messages } = useContext(ChannelContext)

  return (
    <div className="flex flex-col gap-2 h-[90vh] overflow-scroll">
      <ul className="flex gap-5 flex-col-reverse">
        {messages.map((m: Message) => (
          <Message key={m.id} {...m} />
        ))}
      </ul>
    </div>
  )
}
export function Message({
  id,
  content,
  createdAt,
  owner,
  sent,
}: {
  id: string
  content: string
  createdAt: string
  owner: string
  sent: boolean
}) {

  return (
    <div className={`flex flex-row gap-4 ${sent ? '' : 'opacity-50'}`} id={id}>
      <img className="rounded-full" src={`https://vercel.com/api/www/avatar?s=44&teamId=${owner}`} alt={owner} />
      <div className="flex flex-col gap-2">
        <div className="flex flex-row gap-2">
          <p className="font-bold">{owner}</p>
          <p className="text-gray-400">{createdAt}</p>
        </div>
        <h3>{content}</h3>
      </div>
    </div>
  )
}
