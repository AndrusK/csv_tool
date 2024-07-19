<!DOCTYPE html>
<body>

<h1>csv_tool</h1>

<h2>Overview</h2>

<p><code>csv_tool</code> is a command-line utility for processing CSV files. It provides functionalities to find similarities and differences between two CSV files.</p>

<h2>Usage</h2>

<pre><code>./csv_tool -f &lt;firstCSV&gt; -s &lt;secondCSV&gt; -m &lt;mode&gt; -o &lt;outputCSV&gt; [-v] [-h | -help]
</code></pre>

<h3>Options:</h3>

<ul>
<li><code>-f &lt;firstCSV&gt;</code>: Specifies the path to the first CSV file for comparison.</li>
  
<li><code>-s &lt;secondCSV&gt;</code>: Specifies the path to the second CSV file for comparison.</li>
  
<li><code>-m &lt;mode&gt;</code>: Specifies the mode for processing files. Available options:
  <ul>
    <li><code>similarities</code>: Finds similarities between two CSV files.</li>
    <li><code>differences</code>: Finds differences between two CSV files, showing data unique to the second CSV.</li>
  </ul>
</li>

<li><code>-o &lt;outputCSV&gt;</code>: Specifies the output file path where results will be saved. Include the <code>.csv</code> extension.</li>

<li><code>-v</code>: Enables verbose output, providing detailed processing information.</li>

<li><code>-h</code> or <code>-help</code>: Shows this help message, detailing usage and options.</li>
</ul>

<h3>Examples:</h3>

<p>Find similarities between two CSV files:</p>

<pre><code>./csv_tool -f file1.csv -s file2.csv -m similarities -o output.csv
</code></pre>

<p>Find differences (unique data) in the second CSV file:</p>

<pre><code>./csv_tool -f file1.csv -s file2.csv -m differences -o output.csv -v
</code></pre>

</body>
</html>
