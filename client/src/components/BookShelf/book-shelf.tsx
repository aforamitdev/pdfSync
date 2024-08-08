import React from 'react'
import { ResizableHandle, ResizablePanel } from "../ui/resizable"

type Props = {}

const BookShelf = (props: Props) => {
  return (
    <>
          <ResizableHandle withHandle />
          <ResizablePanel defaultSize={32} minSize={30}>
    asas

          </ResizablePanel>

          </>
  )
}

export default BookShelf