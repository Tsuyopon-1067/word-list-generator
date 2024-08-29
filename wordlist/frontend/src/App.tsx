import { useState } from "react";
import { GenerateTable, GenerateCsv } from "../wailsjs/go/main/App";
import { Button, IconButton, TextField, Tooltip } from "@mui/material";
import ContentCopyIcon from "@mui/icons-material/ContentCopy";
import styles from "./App.module.css";
import { HeaderButton } from "./component/HeaderButton";

function App() {
  const [csvText, setCsvText] = useState("");
  const [htmlText, setHtmlText] = useState("");
  const [tooltipTitle, setTooltipTitle] = useState("Copy to Clipboard");

  const generateHtml = () => {
    GenerateTable(csvText).then((res) => setHtmlText(res));
  };

  const generateCsv = () => {
    GenerateCsv(htmlText).then((res) => setCsvText(res));
  };

  const copyToClipboard = async (text: string) => {
    await navigator.clipboard.writeText(text);
    setTooltipTitle("Copied!");
  };

  return (
    <div className={styles.main}>
      <HeaderButton csvText={csvText} setCsvText={setCsvText} htmlText={htmlText} setHtmlText={setHtmlText} />
      <div className={styles.contentArea}>
        <div className={styles.csvArea}>
          <Tooltip title={tooltipTitle} placement="top" arrow>
            <IconButton color="primary" size="small" onClick={() => copyToClipboard(csvText)} onMouseLeave={() => setTooltipTitle("Copy to Clipboard")}>
              <ContentCopyIcon />
            </IconButton>
          </Tooltip>
          <TextField label="csv" multiline fullWidth value={csvText} onChange={(e) => setCsvText(e.target.value)} />
        </div>
        <div className={styles.htmlArea}>
          <Tooltip title={tooltipTitle} placement="top" arrow>
            <IconButton color="primary" size="small" onClick={() => copyToClipboard(htmlText)} onMouseLeave={() => setTooltipTitle("Copy to Clipboard")}>
              <ContentCopyIcon />
            </IconButton>
          </Tooltip>

          <TextField label="html" multiline value={htmlText} fullWidth onChange={(e) => setHtmlText(e.target.value)} />
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
