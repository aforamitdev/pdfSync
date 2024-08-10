import React from 'react'
import { ResizableHandle, ResizablePanel } from '../ui/resizable'

type Props = {}

const Reader = (props: Props) => {
    return (
        <>      <ResizableHandle withHandle />


            <ResizablePanel defaultSize={32} minSize={30}>
                <div className='h-screen'>asas</div>
            </ResizablePanel>
        </>
    )
}

export default Reader