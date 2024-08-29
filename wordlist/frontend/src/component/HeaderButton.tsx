import { useState } from "react";
import React from "react";
import { Save, Open, SaveAs } from "../../wailsjs/go/main/App";
import { Button } from "@mui/material";
import styles from "./HeaderButton.module.css";

type HeaderButtonProps = {
  csvText: string;
  setCsvText: (text: string) => void;
  htmlText: string;
  setHtmlText: (text: string) => void;
};

export const HeaderButton: React.FC<HeaderButtonProps> = ({ csvText, htmlText, setCsvText, setHtmlText }) => {
  const [csvFilePath, setCsvFilePath] = useState("");
  const [htmlFilePath, setHtmlFilePath] = useState("");

  const saveCsv = () => {
    Save(csvText, "csv").then((res) => setCsvFilePath(res));
  };

  const saveHtml = () => {
    Save(htmlText, "html").then((res) => setHtmlFilePath(res));
  };

  const openCsv = () => {
    Open("csv").then((res) => {
      setCsvText(res[0]);
      setCsvFilePath(res[1]);
    });
  };
  const openHtml = () => {
    Open("html").then((res) => {
      setHtmlText(res[0]);
      setHtmlFilePath(res[1]);
    });
  };

  const resave = (content: string, path: string) => {
    if (path === "") {
      return;
    }
    SaveAs(content, path);
  };

  return (
    <div className={styles.header}>
      <div className={styles.header__button}>
        <Button onClick={() => saveCsv()} variant="contained">
          csv保存
        </Button>
        <Button onClick={() => resave(csvText, csvFilePath)} variant="contained">
          csv上書き保存
        </Button>
        <Button onClick={() => openCsv()} variant="contained">
          csv開く
        </Button>
      </div>
      <div className={styles.header__button}>
        <Button onClick={() => saveHtml()} variant="contained">
          html保存
        </Button>
        <Button onClick={() => resave(htmlText, htmlFilePath)} variant="contained">
          html上書き保存
        </Button>
        <Button onClick={() => openHtml()} variant="contained">
          html開く
        </Button>
      </div>
    </div>
  );
};
