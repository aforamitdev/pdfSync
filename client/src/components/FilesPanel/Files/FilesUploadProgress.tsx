import { Progress, Avatar, Box, Card, Flex, Text } from '@radix-ui/themes'
import { Loader2Icon } from 'lucide-react'

type Props = {}

const FilesUploadProgress = (props: Props) => {
    return (
        <>
            <Box className='w-full' >
                <Card>
                    <Flex gap="3" align="center">
                        <Avatar
                            size="3"
                            radius="full"
                            fallback="T"
                        />
                        <Box className='w-full'>
                            <Text as="div" size="2" weight="bold">
                                Teodros Girmay
                            </Text>
                            <Text as="div" size="1" color="gray">
                                450 mbs
                            </Text>
                            <div className='pt-2'>
                                <Progress value={25} size="1" />
                            </div>
                        </Box>
                    </Flex>
                </Card>
            </Box>
        </>
    )
}

export default FilesUploadProgress