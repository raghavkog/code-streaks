const express = require('express');
const app = express();
const port = 8080;

app.use(express.json());

let items = [
    {id: 25, Name: "Raghav"},
    {id: 35, Name: "Manoj"},
];


app.get('/api/items', (req, res) => {
    res.json(items);
  });

app.get('/api/items/:id', (req, res) => {
    const item = items.find(i => i.id === parseInt(req.params.id));
    if (!item) return res.status(404).send('Item not found');
    res.json(item);
  }); 

app.post('/api/items', (req, res) => {
    const newItem = {
        id : req.body.id,
        name: req.body.name,
    };
    items.push(newItem);
    res.status(201).json(newItem);
})

app.put('/api/items/:id', (req, res) => {
    const item = items.find(i => i.id === parseInt(req.params.id));
    if (!item) return res.status(404).send('Item not found');
    item.name = req.body.name;
    res.json(item);
  });
  
app.delete('/api/items/:id', (req, res) => {
    items = items.filter(i => i.id !== parseInt(req.params.id));
    res.status(204).send();
  });
  
app.listen(port, () => {
    console.log(`Server running on http://localhost:${port}`);
  });
  