import { Tabs } from '@radix-ui/react-tabs'
import { ResizableHandle, ResizablePanel } from '../ui/resizable'
import { TabsContent, TabsList, TabsTrigger } from '../ui/tabs'
import { Separator } from '../ui/separator'
import { AlertDescription, AlertTitle } from '../ui/alert'
import { ArrowLeft } from 'lucide-react'

import PdfReader from './PdfReader'
import PageHeader from '../shared/PageHeader'

type Props = {}

const Reader = (props: Props) => {
    return (
        <>
            <ResizableHandle withHandle />
            <ResizablePanel defaultSize={20} maxSize={20} minSize={10}>
                <div className='h-screen'>
                    <div className='  '>

                        <Tabs defaultValue="notes">
                            <div className='mx-2 py-2'>
                                <TabsList className="grid w-full grid-cols-2 ">
                                    <TabsTrigger value="notes">Notes</TabsTrigger>
                                    <TabsTrigger value="activity">Activity</TabsTrigger>
                                </TabsList>
                            </div>
                            <Separator />

                            <TabsContent value="notes">

                                <div className='mx-2'>
                                    <div>
                                        <AlertTitle className='font-medium'>Document Notes</AlertTitle>
                                        <AlertDescription className='text-gray-400 text-xs'>Import Notes To Reffer!</AlertDescription>
                                    </div>
                                    <div>ONe one </div>
                                </div>
                            </TabsContent>

                            <TabsContent value="activity">
                                <div className='mx-2'>
                                    <div>
                                        <AlertTitle className='font-medium'>Document Activity!</AlertTitle>
                                        <AlertDescription className='text-gray-400 text-xs'>Recent activity on the document </AlertDescription>

                                    </div>
                                </div>
                            </TabsContent>


                        </Tabs>

                    </div>

                </div>
            </ResizablePanel>
            <ResizableHandle withHandle />
            <ResizablePanel defaultSize={60} minSize={30}>
                <div className='h-screen '>
                    <PageHeader title='Document Reader' subTitle='' />
                    <div>
                        <PdfReader />
                    </div>
                </div>
            </ResizablePanel>
        </>
    )
}

export default Reader