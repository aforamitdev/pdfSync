import { useEffect } from 'react';
import { useDropzone } from 'react-dropzone';

type Props = {
    setFile: React.Dispatch<React.SetStateAction<File | null>>
}

function FileUpload({ setFile }: Props) {
    const { acceptedFiles, getRootProps, getInputProps } = useDropzone();

    useEffect(() => {
        if (acceptedFiles.length) {
            setFile(acceptedFiles[0])
        }
    }, [acceptedFiles])

    return (
        <section className="border w-full flex justify-center border-dashed py-5 my-2">
            <div {...getRootProps({ className: 'dropzone' })}>
                <input {...getInputProps()} />
                <p>Select a file to upload ...</p>
            </div>

        </section>
    )
}

export default FileUpload