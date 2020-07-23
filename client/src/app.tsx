import { Window, View, Text, Image, hot } from "@nodegui/react-nodegui";
import React from "react";
import { MemoryRouter } from "react-router";
import AppRoutes from "./routes";
import { wallpaperDataUri } from "./styles/wallpaper";
const minSize = { width: 1920, height: 1080 };

class App extends React.Component {
  render() {
    return (
      <MemoryRouter>
        <Window
          windowTitle="D2 Mod Maker"
          minSize={minSize}
          style={containerStyle}
        >
          <Text>
            {`<img src="${wallpaperDataUri}"></img>`}
          </Text>
          <View style={containerStyle}>
            <AppRoutes />
          </View>
        </Window>
      </MemoryRouter >
    );
  }
}

const containerStyle = `
  flex: 1; 
  width: 1920px;
  height: 1080px;
  background-image: url("https://images8.alphacoders.com/613/613547.jpg");
`


export default hot(App);
