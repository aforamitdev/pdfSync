import { ResizableHandle, ResizablePanel } from '../ui/resizable'

type Props = {}

const Reader = (props: Props) => {
    return (
        <>
            <ResizableHandle withHandle />
            <ResizablePanel defaultSize={20} minSize={10}>
                <div className='h-screen'>asas</div>
            </ResizablePanel>
            <ResizableHandle withHandle />
            <ResizablePanel defaultSize={60} minSize={30}>
                <div className='h-screen'>asas</div>
            </ResizablePanel>
        </>
    )
}

export default Reader