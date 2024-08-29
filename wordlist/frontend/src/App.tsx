import { useState } from "react";
import { GenerateTable, GenerateCsv } from "../wailsjs/go/main/App";
import { Button, TextField } from "@mui/material";
import styles from "./App.module.css";
import { HeaderButton } from "./component/HeaderButton";

function App() {
  const [csvText, setCsvText] = useState("");
  const [htmlText, setHtmlText] = useState("");
  const generateHtml = () => {
    GenerateTable(csvText).then((res) => setHtmlText(res));
  };

  const generateCsv = () => {
    GenerateCsv(htmlText).then((res) => setCsvText(res));
  };

  return (
    <div className={styles.main}>
      <HeaderButton csvText={csvText} setCsvText={setCsvText} htmlText={htmlText} setHtmlText={setHtmlText} />
      <div className={styles.contentArea}>
        <div className={styles.csvArea}>
          <TextField label="Multiline" multiline fullWidth value={csvText} onChange={(e) => setCsvText(e.target.value)} />
        </div>
        <div className={styles.htmlArea}>
          <TextField label="Multiline" multiline value={htmlText} fullWidth onChange={(e) => setHtmlText(e.target.value)} />
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
