import okwolo from 'okwolo/lite';

import './index.scss';

const app = okwolo(document.body);

app.setState({});

app(() => () => (
    ['h2', {a: 3}, [
        'targetblank',
    ]]
));
