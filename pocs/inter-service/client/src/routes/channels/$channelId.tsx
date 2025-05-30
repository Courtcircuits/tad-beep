import { createFileRoute } from '@tanstack/react-router'
import { ChannelProvider } from '../../context/channel'
import { MessageList } from '../../components/message-list'
import { MessageInput } from '../../components/message-input'
import { Search } from '@/components/search'

export const Route = createFileRoute('/channels/$channelId')({
  component: RouteComponent,
})

function RouteComponent() {
  const { channelId: channelId } = Route.useParams()
  return (
    <div>
      <p>Hello on the channel {channelId}, Press ctrl+j to open the search bar</p>
      <div className="mx-[20px] flex flex-col justify-between gap-2 ">
        <ChannelProvider channelID={channelId}>
          <MessageList />
          <MessageInput />
          <Search/>
        </ChannelProvider>
      </div>
    </div>
  )
}
