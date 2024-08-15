import React from 'react'
import { ArrowLeft } from 'lucide-react'
import { Separator } from '../ui/separator'
import { Flex, Heading, Text } from '@radix-ui/themes'

type Props = {
    title: string
    subTitle: string
}

const PageHeader = ({ title }: Props): React.ReactElement => {
    return (
        <>
            <div className='flex items-center py-3.5 mx-3'>
                <div className='px-2 border  py-1 rounded-md shadow-sm '>
                    <ArrowLeft size={18} />
                </div>
                <Flex className='flex flex-col'>
                    <Heading size="4" className='text-sm font-bold mx-2'>{title}</Heading>
                </Flex>

            </div>
            <Separator />
        </>
    )
}

export default PageHeader