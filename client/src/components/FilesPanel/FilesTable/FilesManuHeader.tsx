import { Box, Flex, IconButton, Select, TextField } from '@radix-ui/themes'
import { File, SearchIcon } from 'lucide-react'
import React from 'react'

type Props = {}

const FilesManuHeader = (props: Props) => {
    const [value, setValue] = React.useState('light');

    return (
        <div className='flex   justify-between '>

            <div className='flex w-1/2 gap-x-2'>
                <div className='pb-3 w-full'>
                    <Flex direction="column" >
                        <Select.Root value={value} onValueChange={setValue}>
                            <Select.Trigger>
                                <Flex as="span" align="center" gap="2">
                                    <File size="15" />
                                </Flex>
                            </Select.Trigger>
                            <Select.Content position="popper">
                                <Select.Item value="light">Light</Select.Item>
                                <Select.Item value="dark">Dark</Select.Item>
                            </Select.Content>
                        </Select.Root>
                    </Flex>
                </div>
                <div className='pb-3 w-full'>
                    <Flex direction="column" >
                        <Select.Root value={value} onValueChange={setValue}>
                            <Select.Trigger>
                                <Flex as="span" align="center" gap="2">
                                    <File size="15" />
                                </Flex>
                            </Select.Trigger>
                            <Select.Content position="popper">

                                <Select.Item value="light">Light</Select.Item>
                                <Select.Item value="dark">Dark</Select.Item>
                            </Select.Content>
                        </Select.Root>
                    </Flex>
                </div>

                <div className='pb-3 w-full'>
                    <Flex direction="column" >
                        <Select.Root value={value} onValueChange={setValue}>
                            <Select.Trigger>
                                <Flex as="span" align="center" gap="2">
                                    <File size="15" />
                                </Flex>
                            </Select.Trigger>
                            <Select.Content position="popper">

                                <Select.Item value="light">Light</Select.Item>
                                <Select.Item value="dark">Dark</Select.Item>
                            </Select.Content>
                        </Select.Root>
                    </Flex>
                </div>

            </div>
            <div>
                <Box maxWidth="250px">
                    <TextField.Root placeholder="Search the docs…" size="2">
                        <TextField.Slot>
                            <SearchIcon height="16" width="16" />
                        </TextField.Slot>
                        <TextField.Slot>

                        </TextField.Slot>
                    </TextField.Root>
                </Box>
            </div>




        </div>
    )
}

export default FilesManuHeader