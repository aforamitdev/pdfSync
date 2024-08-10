
import { CardDescription, CardTitle } from "../ui/card"
import { Input } from "../ui/input"
import { ResizableHandle, ResizablePanel } from "../ui/resizable"
import { Separator } from "../ui/separator"
import { FileList } from "./FileLists"
import FileProgress from "./FileProgress"
type Props = {}

const BookShelf = (props: Props) => {
  return (
    <>
      <ResizableHandle withHandle />
      <ResizablePanel defaultSize={32} minSize={30}>
        <div className="h-screen px-4 py-4 " >

          <div className="flex items-center justify-between space-y-2 pb-2">
            <div>
              <CardTitle>Files and Assets!</CardTitle>
              <CardDescription className="text-muted-foreground">
                Ebooks and Pdfs Uploaded to this system.
              </CardDescription>
            </div>
            <div className="flex items-center space-x-2">
              {/* <UserNav /> n        */}
            </div>
          </div>
          <Separator />
          {/* file upload box */}
          <div className="grid w-full  items-center gap-1.5 h-20  my-2">
            <Input id="asset" type="file" className="h-20 flex  item-center" />
          </div>

          {/* progress bar */}

          <FileProgress />






          <FileList />
        </div >
      </ResizablePanel>

    </>
  )
}

export default BookShelf