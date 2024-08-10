import React from 'react'
import { Progress } from '../ui/progress'
import { File, Terminal, Loader2Icon } from 'lucide-react'
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar'

type Props = {}

function FileProgress({ }: Props) {

    const [progress, setProgress] = React.useState(13)

    React.useEffect(() => {
        const timer = setTimeout(() => setProgress(66), 500)
        return () => clearTimeout(timer)
    }, [])

    return <div className="flex  justify-items items-center gap-4 text-sm border rounded-md py-2.5 px-4">
        <div >

            <Avatar>
                <AvatarImage alt={"AS"} />
                <AvatarFallback>
                    SA
                </AvatarFallback>
            </Avatar>
        </div>
        <div className="grid gap-1 flex-1 ">
            <div className="font-semibold text-sm flex justify-between items-center">
                <div>Programming with Linux.pdf</div>
                <div><Loader2Icon className=' h-4' /></div>
            </div>
            <div className="line-clamp-1 text-xs text-gray-700">230KB</div>
            <div className="line-clamp-1 text-xs w-full ">
                <Progress value={33} className='h-2' />
            </div>
        </div>
    </div>

}

export default FileProgress