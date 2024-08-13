import { Tabs } from '@radix-ui/react-tabs'
import { ResizableHandle, ResizablePanel } from '../ui/resizable'
import { TabsContent, TabsList, TabsTrigger } from '../ui/tabs'
import { Card, CardHeader } from '../ui/card'
import { Separator } from '../ui/separator'
import { Alert, AlertDescription, AlertTitle } from '../ui/alert'
import { ArrowLeft, ChevronLeft } from 'lucide-react'
import {
    AreaHighlight,
    Highlight,
    PdfHighlighter,
    PdfLoader,
    Popup,
    Tip,
} from "react-pdf-highlighter";
import PdfReader from './PdfReader'

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
                    <div className='flex items-center py-3.5 mx-3'>
                        <div className='px-2 border  py-1 rounded-md shadow-sm '>
                            <ArrowLeft size={18} />
                        </div>
                        <h4 className='text-sm font-bold mx-2'>Programmign with linux.pdf</h4>
                    </div>
                    <Separator />
                    <div className='mx-2'>

                        <PdfReader />
                    </div>
                </div>
            </ResizablePanel>
        </>
    )
}

export default Reader