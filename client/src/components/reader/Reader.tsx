import { Tabs } from '@radix-ui/react-tabs'
import { ResizableHandle, ResizablePanel } from '../ui/resizable'
import { TabsContent, TabsList, TabsTrigger } from '../ui/tabs'
import { Card, CardHeader } from '../ui/card'
import { Separator } from '../ui/separator'
import { Alert, AlertDescription, AlertTitle } from '../ui/alert'

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
                                    </div>
                                    <div>ONe one </div>
                                </div>
                            </TabsContent>

                            <TabsContent value="activity">
                                <div className='mx-2'>
                                    <div>
                                        <AlertTitle className='font-medium'>Document Activity!</AlertTitle>
                                    </div>
                                </div>
                            </TabsContent>


                        </Tabs>

                    </div>

                </div>
            </ResizablePanel>
            <ResizableHandle withHandle />
            <ResizablePanel defaultSize={60} minSize={30}>
                <div className='h-screen'>
                    PDFS
                </div>
            </ResizablePanel>
        </>
    )
}

export default Reader