//App component bootstraps the entire react app
class App extends React.Component {
    render() {
        //we would put auth component decisions here when needed
        return <TypingTest />
    }
}

class TypingTest extends React.Component {
    render() {
        const typingTestUrl = '/start';
        return (
            <div className="container">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>Typing Test</h1>
                    <p>test test test</p>
                    <p>test test test</p>
                    <a href={typingTestUrl} className="btn btn-primary btn-lg btn-login btn-block">Take test</a>
                </div>
            </div>
        )
    }
}

//tells react where to render app
ReactDOM.render(<App />, document.getElementById('app'));