import { Document, Page, pdfjs } from "react-pdf"
import pdf from "../../examples/concurrencyingo.pdf"
import { useState } from "react"
import * as pdfJS from "pdfjs-dist";
import pdfJSWorkerURL from "pdfjs-dist/build/pdf.worker?url";
pdfJS.GlobalWorkerOptions.workerSrc = pdfJSWorkerURL;
type Props = {}
// https://codesandbox.io/s/react-pdf-sample-forked-displaying-pdf-using-react-dyqr51?file=/src/components/pdf/single-page.js:640-643
const PdfReader = (props: Props) => {
    const [s, setS] = useState(pdf)
    return (
        <div>

            <input
                type="file"
                onChange={(e) => {
                    const file = e.target.files[0];
                    const reader = new FileReader();
                    reader.readAsDataURL(file);
                    reader.onload = () => {
                        setS(reader.result);
                    };
                }} />
            <Document file={s} onError={e => {
                console.log(e)
            }}

                onLoadSuccess={e => {
                    console.log(e, "first")
                }}
                onLoad={e => {
                    console.log(e, "LOAD")
                }}
            >
                <Page pageNumber={1} />
            </Document>

        </div>
    )
}

export default PdfReader