import React, { useState } from 'react';
import './style.scss';

function numberWithCommas(x) {
  return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',');
}

const ws = new WebSocket(`${(window.location.protocol === 'https:' ? 'wss://' : 'ws://') + window.location.host}/ws`);
ws.onopen = () => console.log('connected');
ws.onclose = () => console.log('disconnected');

const Content = () => {
  const [bitCoins, setBitCoins] = useState([]);
  ws.onmessage = evt => {
    const message = JSON.parse(evt.data);
    setBitCoins([...message]);
  };

  return (
    <main>
      <table>
        <thead>
          <tr>
            <th>來源</th>
            <th>價格</th>
            <th>24小時交易量</th>
            <th>24小時漲跌幅度</th>
            <th>總市值</th>
          </tr>
        </thead>
        <tbody>
          {bitCoins.map(bitcoin => (
            <tr key={bitcoin.source_name}>
              <td>{bitcoin.source_name}</td>
              <td>{`US$${numberWithCommas(Math.round(bitcoin.price))}`}</td>
              <td>{`US$${numberWithCommas(Math.round(bitcoin.volume_24h))}`}</td>
              <td>{`${Math.round(bitcoin.percent_change_24h * 100) / 100}%`}</td>
              <td>{`US$${numberWithCommas(Math.round(bitcoin.market_cap))}`}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </main>
  );
};

export default Content;
