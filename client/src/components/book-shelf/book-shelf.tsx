
import FileUpload from "../shared/FileUpload"
import PageHeader from "../shared/PageHeader"
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
        <div className="h-screen " >
          <PageHeader title="Library" subTitle="sasas" />
          {/* file upload box */}
          <div className="bg-gray-100 h-full py-2 px-10">

            <div className="bg-white">
              <FileUpload setFile={() => { }} />
            </div>
            <FileProgress />

            <FileList />
          </div>
        </div >
      </ResizablePanel>

    </>
  )
}

export default BookShelf