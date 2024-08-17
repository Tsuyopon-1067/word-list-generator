import { useState } from "react";
import { GenerateTable, GenerateCsv, Save, Open } from "../wailsjs/go/main/App";
import { Button, TextField } from "@mui/material";
import styles from "./App.module.css";

function App() {
  const [csvText, setCsvText] = useState("");
  const [htmlText, setHtmlText] = useState("");
  const [fileName, setFileName] = useState("");
  const generateHtml = () => {
    GenerateTable(csvText).then((res) => setHtmlText(res));
  };

  const generateCsv = () => {
    GenerateCsv(htmlText).then((res) => setCsvText(res));
  };

  const saveFile = (content: string, ext: string) => {
    Save(content, ext);
  }

  const openFile = (ext: string, setState: (s: string) => void) => {
    Open(ext).then((res) => setState(res));
  }

  return (
    <div className={styles.main}>
      <div className={styles.titleArea}>
        <Button
          onClick={() => {
            saveFile(csvText, 'csv')
          }}
          variant="contained"
        >
          csv保存
        </Button>
        <Button
          onClick={() => {
            saveFile(csvText, 'html')
          }}
          variant="contained"
        >
          html保存
        </Button>

        <Button
          onClick={() => {
            openFile('csv', setCsvText)
          }}
          variant="contained"
        >
          csv開く
        </Button>
        <Button
          onClick={() => {
            openFile('html', setHtmlText)
          }}
          variant="contained"
        >
          html開く
        </Button>

        <Button variant="contained">
          csvコピー
        </Button>
        <Button variant="contained">
          htmlコピー
        </Button>
      </div>
      <div className={styles.contentArea}>
        <div className={styles.csvArea}>
          <TextField
            label="Multiline"
            multiline
            fullWidth
            value={csvText}
            onChange={(e) => setCsvText(e.target.value)}
          />
        </div>
        <div className={styles.htmlArea}>
          <TextField
            label="Multiline"
            multiline
            value={htmlText}
            fullWidth
            onChange={(e) => setHtmlText(e.target.value)}
          />
        </div>
        <div className={styles.csvToHtmlButtonArea}>
          <Button onClick={generateHtml} variant="contained">
            →
          </Button>
        </div>
        <div className={styles.htmlToCsvButtonArea}>
          <Button onClick={generateCsv} variant="contained">
            ←
          </Button>
        </div>
      </div>
    </div>
  );
}

export default App;
