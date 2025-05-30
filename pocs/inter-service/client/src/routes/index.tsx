import { createFileRoute } from '@tanstack/react-router'
import { JoinForm } from '../components/join-form'

export const Route: unknown = createFileRoute('/')({
  component: Index,
})

function Index() {
  return (
    <div className="p-2">
      <h3>👋 Welcome to Beep<sup>experimental</sup></h3>
      <h2>This a proof of concept for inter-service communication leveraging :</h2>
      <ul>
        <li>🔗 gRPC</li>
        <li>🔗 Quickwit</li>
        <li>🔗 GrapQL</li>
      </ul>
      <JoinForm/>
    </div>
  )

}
