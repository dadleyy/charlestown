<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" version="1.0">
  <xsl:param name="version"/>
  <xsl:output method="xml" indent="yes"/>
  <xsl:template match="plist-template">
    <plist version="1.0">
      <dict>
        <key>CFBundleDisplayName</key>
        <string>Charlestown</string>
        <key>CFBundleDisplayName</key>
        <string>Charlestown</string>
        <key>CFBundleIdentifier</key>
        <string>com.charlestown.terminal</string>
        <key>CFBundlePackageType</key>
        <string>APPL</string>
        <key>CFBundleSignature</key>
        <string>cctw</string>
        <key>CFBundleExecutable</key>
        <string>charlestown</string>
        <key>NSHighResolutionCapable</key>
        <true/>
        <key>CFBundleIconFile</key>
        <string>charlestown-icon</string>
        <key>CFBundleVersion</key>
        <string><xsl:value-of select="$version"/></string>
      </dict>
    </plist>
  </xsl:template>
</xsl:stylesheet>
