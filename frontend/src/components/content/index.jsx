import React, { useState, useEffect } from 'react';
import './style.scss';

function numberWithCommas(x) {
  return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ',');
}

const Content = () => {
  const [bitCoins, setBitCoins] = useState([]);
  useEffect(() => {
    fetch('/api/btc')
      .then(response => response.json())
      .then(result => setBitCoins([...result]))
      .catch(error => error);
  }, []);

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
