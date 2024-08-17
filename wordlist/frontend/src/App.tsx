import { useState } from "react";
import { GenerateTable, GenerateCsv, Save, Open, SaveAs } from "../wailsjs/go/main/App";
import { Button, TextField } from "@mui/material";
import styles from "./App.module.css";

function App() {
  const [csvText, setCsvText] = useState("");
  const [htmlText, setHtmlText] = useState("");
  const [csvFilePath, setCsvFilePath] = useState("");
  const [htmlFilePath, setHtmlFilePath] = useState("");
  const generateHtml = () => {
    GenerateTable(csvText).then((res) => setHtmlText(res));
  };

  const generateCsv = () => {
    GenerateCsv(htmlText).then((res) => setCsvText(res));
  };

  const saveCsv = () => {
    Save(csvText, 'csv').then((res) => setCsvFilePath(res));
  }

  const saveHtml = () => {
    Save(htmlText, 'html').then((res) => setHtmlFilePath(res));
  }

  const openCsv = () => {
    Open('csv')
      .then((res) => {
        setCsvText(res[0]);
        setCsvFilePath(res[1]);
      });
  }
  const openHtml = () => {
    Open('html')
      .then((res) => {
        setHtmlText(res[0]);
        setHtmlFilePath(res[1]);
      });
  }


  const resave = (content: string, path: string) => {
    if (path === "") {
      return;
    }
    SaveAs(content, path);
  }

  return (
    <div className={styles.main}>
      <div className={styles.titleArea}>
        <Button
          onClick={() => saveCsv()}
          variant="contained"
        >
          csv保存
        </Button>
        <Button
          onClick={() => saveHtml()}
          variant="contained"
        >
          html保存
        </Button>

        <Button
          onClick={() => resave(csvText, csvFilePath)}
          variant="contained"
        >
          csv上書き保存
        </Button>
        <Button
          onClick={() => resave(htmlText, htmlFilePath)}
          variant="contained"
        >
          html上書き保存
        </Button>

        <Button
          onClick={() => openCsv()}
          variant="contained"
        >
          csv開く
        </Button>
        <Button
          onClick={() => openHtml()}
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
