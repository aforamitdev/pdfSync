import React from 'react'
import { Progress } from '../ui/progress'
import { File, Terminal, Loader2Icon } from 'lucide-react'
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar'
import FilesItem from './Files/FilesItem'

type Props = {}

function FileProgress({ }: Props) {

    const [progress, setProgress] = React.useState(13)

    React.useEffect(() => {
        const timer = setTimeout(() => setProgress(66), 500)
        return () => clearTimeout(timer)
    }, [])

    return <div className="flex  flex-col justify-items items-center gap-4 text-sm  ">

        <FilesItem />
        <FilesItem />


    </div>

}

export default FileProgress