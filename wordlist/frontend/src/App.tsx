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

  // ファイル読み込みボタン
  const FolderSelector = () => {
    const handleFolderSelect = (event: any) => {
      const file = event.target.files[0];
      const reader = new FileReader();
      reader.onload = () => {
        if (reader.result) {
          const filename = file.name;
          const ext = filename.split('.').pop();
          const content = reader.result?.toString()
          setFileName(filename);
          if (ext == 'csv') {
            setCsvText(content)
          } else if (ext == 'html') {
            setHtmlText(content)
          } else {
            // TODO: トーストを出す
          }
        }
      }
      reader.readAsText(file);
    };

    return (
      <span>
        <input
          type="file"
          style={{ display: "none" }}
          onChange={handleFolderSelect}
          id="folder-input"
          accept='.csv, .html'
        />
        <label htmlFor="folder-input">
          <Button variant="contained" component="span">
            フォルダを選択
          </Button>
        </label>
      </span>
    );
  };

  return (
    <div className={styles.main}>
      <div className={styles.titleArea}>
        <FolderSelector />
        <span style={{ color: "black" }}> filename: {fileName}</span>

        <Button onClick={() => {
          saveFile(csvText, 'csv')
        }}>csv保存</Button>
        <Button onClick={() => {
          saveFile(csvText, 'html')
        }}>html保存</Button>

        <Button onClick={() => {
          openFile('csv', setCsvText)
        }}>csv開く</Button>
        <Button onClick={() => {
          openFile('html', setHtmlText)
        }}>html開く</Button>

        <Button>csvコピー</Button>
        <Button>htmlコピー</Button>
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
