import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { Progress } from '@/components/ui/progress'
import { Loader2Icon } from 'lucide-react'

type Props = {}

const FilesItem = (props: Props) => {
    return (
        <>
            <div className='border rounded-md flex py-2.5  w-full'>
                <div className='px-2'>
                    <Avatar>
                        <AvatarImage alt={"AS"} />
                        <AvatarFallback>
                            SA
                        </AvatarFallback>
                    </Avatar>
                </div>
                <div className="grid gap-1 flex-1 ">
                    <div>

                        <div className="font-semibold text-sm flex justify-between items-center">
                            <div>Programming with Linux.pdf</div>
                            <div className='px-2'><Loader2Icon className=' h-4' /></div>
                        </div>
                        <div className="line-clamp-1 text-xs text-gray-700">230KB</div>
                        <div className="line-clamp-1 text-xs w-full pr-2">
                            <Progress value={33} className='h-2' />
                        </div>
                    </div>


                </div>


            </div >
        </>
    )
}

export default FilesItem