import { useState } from 'react';
import { GenerateTable } from '../wailsjs/go/main/App';
import { Box } from '@mui/system';
import { Button, ButtonBase, TextField } from '@mui/material';
import styles from './App.module.css';

function App() {
    const [csvText, setCsvText] = useState('');
    const [htmlText, setHtmlText] = useState('');
    const generateHtml = () => {
        GenerateTable(csvText).then((res) => setHtmlText(res));
    }

    const generateCsv = () => {
        //GenerateTable('');
    }

    return (
        <div className={styles.main}>
            <div className={styles.titleArea}>
                <Button variant="contained">open file</Button>
                <Button variant="contained">prev history</Button>
                <Button variant="contained">next history</Button>
            </div>
            <div className={styles.contentArea}>
                <div className={styles.csvArea}>
                    <TextField
                        label='Multiline'
                        multiline
                        fullWidth
                        value={csvText}
                        onChange={
                            (e) => setCsvText(e.target.value)
                        }
                    />
                </div>
                <div className={styles.htmlArea}>
                    <TextField
                        label='Multiline'
                        multiline
                        value={htmlText}
                        fullWidth
                        onChange={
                            (e) => setHtmlText(e.target.value)
                        }
                    />
                </div>
                <div className={styles.csvToHtmlButtonArea}>
                    <Button onClick={generateHtml} variant="contained">→</Button>
                </div>
                <div className={styles.htmlToCsvButtonArea}>
                    <Button onClick={generateCsv} variant="contained">←</Button>
                </div>
            </div>
        </div>
    )
}

export default App
