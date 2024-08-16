import React from 'react'
import FileUploadProgress from './Files/FilesUploadProgress'

type Props = {}

function FileProgress({ }: Props) {

    const [progress, setProgress] = React.useState(13)

    React.useEffect(() => {
        const timer = setTimeout(() => setProgress(66), 500)
        return () => clearTimeout(timer)
    }, [])

    return <div className="flex  flex-col justify-items items-center gap-1 text-sm  ">
        <FileUploadProgress />
        <FileUploadProgress />
    </div>

}

export default FileProgress