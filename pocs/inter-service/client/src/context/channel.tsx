import { createContext, useState } from 'react'
import { v4 as uuidv4 } from 'uuid'
import { useSubscription, useMutation,useLazyQuery } from '@apollo/client'
import { gql } from '../__generated__/gql'
import type { NewMessage } from '../__generated__/graphql'

export interface Message {
  id: string
  content: string
  createdAt: string
  owner: string
  sent: boolean
}

interface ChannelContextType {
  channelId: string
  messages: Message[]
  sendMessage: (message: NewMessage) => void
  search(query: string): void
  searchResults: Message[]
}


export const ChannelContext = createContext<ChannelContextType>({
  channelId: '',
  messages: [],
  sendMessage: (_message: NewMessage) => { },
  search: (_query: string) => { },
  searchResults: [],
})

const SEND_MESSAGES = gql(`
mutation SendMessage($message: NewMessage!) {
    sendMessage(message: $message) {
        id
        content
        createdAt
        owner
    }
}
`)

const SEARCH_MESSAGES = gql(`
query SearchMessages($query: String!, $channelID: String!) {
    searchMessages(query:$query, channelID:$channelID) {
        id
        content
        createdAt
        owner
        channelID
    }
}
`)

const SUBSCRIBE_MESSAGES = gql(`
subscription GetMessages($channelID: String!, $ownerID: String!) {
    getMessages(channelID: $channelID, ownerID: $ownerID) {
        id
        content
        createdAt
        owner
    }
}
`)
export function ChannelProvider({ children, channelID }: { channelID: string, children: React.ReactNode }) {

  const [skeletons, setSkeletons] = useState<string[]>([])
  const [sendMessage, { }] = useMutation(SEND_MESSAGES, {
    onCompleted(data, _clientOptions) {
      if (data.sendMessage) {
        const message = data.sendMessage
        if (!message) return
        const oldMessages = (oldMessage: Message[]) => oldMessage.filter((m) => {
          for (const skeleton of skeletons) {
            if (m.id === skeleton) return false
          }
          return true
        })
        setMessages((messages) => [...oldMessages(messages), {
          ...message,
          sent: true,
        }])
        setSkeletons([])
      }
    },
  });
  const [messages, setMessages] = useState<Message[]>([])

  const { } = useSubscription(
    SUBSCRIBE_MESSAGES,
    {
      variables: {
        channelID,
        ownerID: localStorage.getItem('name') || '',
      },
      onData: ({ data }) => {
        if (data.data?.getMessages) {
          const message = data.data?.getMessages[0]
          if (!message) return
          setMessages((messages) => [...messages, {
            ...message,
            sent: true,
          }])
        }
      },
    }
  );

  const [refetchSearch, { data  }] = useLazyQuery(SEARCH_MESSAGES);

  const addMessage = (message: NewMessage) => {
    const uuid = uuidv4()
    setSkeletons((skeletons) => [...skeletons, uuid])

    setMessages((messages) => [...messages, {
      content: message.content,
      id: uuid,
      owner: message.owner,
      createdAt: new Date().toISOString(),
      sent: false,
    }])
  }


  const messages_ordered = messages.sort((a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime())

  console.log(data)

  const message_results =data?.searchMessages.map((message) => ({
    ...message,
    sent: true,
  })) || []


  return (
    <ChannelContext.Provider value={
      {
        channelId: channelID, messages: messages_ordered, sendMessage: (message) => {
          addMessage(message)
          sendMessage({ variables: { message } })
        },
        search: (query: string) => {
          refetchSearch({variables: {query, channelID}})
        },
        searchResults: message_results,
      }
    }
      >
      { children }
    </ ChannelContext.Provider>
      )
}
