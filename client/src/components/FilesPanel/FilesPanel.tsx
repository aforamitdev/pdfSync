
import FileUpload from "../shared/FileUpload"
import PageHeader from "../shared/PageHeader"
import { ResizableHandle, ResizablePanel } from "../ui/resizable"
import FileProgress from "./FileProgress"
import FilesTable from "./FilesTable/FilesTable"
import { columns } from "./FilesTable/TableColumns"
import { TFilesList } from "./types"

const data: TFilesList[] = [
  {
    id: 1,
    file: "Work station ",
    creator: "Amit rai",
    createdat: new Date(),
    access: "developers"


  },
  {
    id: 2,
    file: "Project init",
    creator: "Sidhi Rai",
    createdat: new Date(),
    access: "Product Manager"


  },
  {
    id: 3,
    file: "Transings files",
    creator: "Raman ",
    createdat: new Date(),
    access: "DEV"


  },
  {
    id: 4,
    file: "receip",
    creator: "Mickal jekson", createdat: new Date(),
    access: "Account Admin"


  },
  {
    id: 5,
    file: "Onboargin camping ",
    creator: "Eax",
    createdat: new Date(),
    access: "HR Admin"

  },
]


const FilesPanel = () => {
  return (
    <>
      <ResizableHandle withHandle />
      <ResizablePanel defaultSize={32} minSize={30}>
        <div className="h-screen " >
          <PageHeader title="Files and assets" subTitle="sasas" />
          {/* file upload box */}
          <div className="bg-gray-100 h-full py-2 px-10">
            <div className="bg-white">
              <FileUpload setFile={() => { }} />
            </div>

            <FileProgress />

            {/* <FileList /> */}
            <FilesTable data={data} columns={columns} />
          </div>
        </div >
      </ResizablePanel>

    </>
  )
}

export default FilesPanel