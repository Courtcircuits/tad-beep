import { useContext, useEffect, useState } from "react"
import {
  CommandDialog,
  CommandEmpty,
  CommandInput,
} from "@/components/ui/command"
import { ChannelContext } from "@/context/channel"


export function Search() {
  const [open, setOpen] = useState(false)
  const [value, setValue] = useState("")
  const {search, searchResults} = useContext(ChannelContext)

  useEffect(() => {
    const down = (e: KeyboardEvent) => {
      if (e.key === "j" && (e.metaKey || e.ctrlKey)) {
        e.preventDefault()
        setOpen((open) => !open)
      }
    }

    document.addEventListener("keydown", down)
    return () => document.removeEventListener("keydown", down)
  }, [])
  useEffect(() => {
    if (value.length > 0) {
      search(value)
    }
  }, [value])
  console.log(searchResults)
  return (
    <CommandDialog open={open} onOpenChange={setOpen} className="bg-gray-500">
      <CommandInput placeholder="Type a command or search..." onValueChange={setValue} />
      {
        searchResults.length === 0  ? <CommandEmpty>No results found.</CommandEmpty> :
          searchResults.map((item, index) => (
            <p key={index}>{item.content}</p>
          ))
      }
    </CommandDialog>
  )
}
