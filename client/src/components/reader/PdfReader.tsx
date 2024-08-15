import React from 'react'
import { Tabs } from '@radix-ui/themes'
import { Content, List, Trigger } from '@radix-ui/react-tabs'
type Props = {}

const PdfReader = (props: Props) => {
    return (
        <div className='bg-gray-100 h-screen'>
            <Tabs.Root defaultValue="account">
                <Tabs.List className='bg-white'>
                    <Tabs.Trigger value="account">Account</Tabs.Trigger>
                    <Tabs.Trigger value="setting">setting</Tabs.Trigger>
                </Tabs.List>

                <Tabs.Content value='account'>
                    <div className='w-full flex justify-center h-screen '>
                        <div className='bg-white w-3/4 my-5 h-full'>account</div>
                    </div>
                </Tabs.Content>
                <Tabs.Content value='setting'>
                    <div className='w-full flex justify-center h-screen '>
                        <div className='bg-white w-3/4 my-5 h-full'>setting</div>
                    </div>
                </Tabs.Content>
            </Tabs.Root>


        </div>
    )
}

export default PdfReader