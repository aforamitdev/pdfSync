import { Document, Page } from "react-pdf"
import pdf from "../../examples/concurrencyingo.pdf"
type Props = {}
// https://codesandbox.io/s/react-pdf-sample-forked-displaying-pdf-using-react-dyqr51?file=/src/components/pdf/single-page.js:640-643
const PdfReader = (props: Props) => {
    return (
        <div>
            <Document file={"https://pdfobject.com/pdf/sample.pdf"} onError={e => {
                console.log(e)
            }}
                options={{}}
            >
                <Page pageNumber={1} />
            </Document>

        </div>
    )
}

export default PdfReader